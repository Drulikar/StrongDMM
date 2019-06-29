package io.github.spair.strongdmm.logic.dme

import io.github.spair.strongdmm.common.*

class DmeItem(
    private val environment: Dme,
    val type: String,
    val vars: Map<String, String>,
    val children: List<String>
) {

    private val lookedVars = mutableMapOf<String, String?>()

    private val parentType by lazy {
        when (type) {
            TYPE_DATUM -> null
            TYPE_ATOM -> TYPE_DATUM
            TYPE_ATOM_MOVABLE, TYPE_AREA, TYPE_TURF -> TYPE_ATOM
            TYPE_OBJ, TYPE_MOB -> TYPE_ATOM_MOVABLE
            else -> type.substringBeforeLast('/')
        }
    }

    fun getParent(): DmeItem? = parentType?.let { environment.getItem(it) }

    fun getVar(name: String): String? {
        if (vars.containsKey(name)) {
            return vars[name]
        } else if (lookedVars.containsKey(name)) {
            return lookedVars[name]
        }

        return parentType?.let {
            val parent = environment.getItem(it)!!
            parent.getVar(name).apply { lookedVars[name] = this }
        }
    }

    fun getVarText(name: String) = getVar(name)?.takeIf { it.isNotEmpty() }?.run { substring(1, length - 1) }

    fun getVarFloat(name: String): Float {
        return try {
            getVar(name)?.toFloat() ?: NON_EXISTENT_FLOAT
        } catch (e: NumberFormatException) {
            NON_EXISTENT_FLOAT
        }
    }

    fun getVarInt(name: String): Int {
        return try {
            getVar(name)?.toInt() ?: NON_EXISTENT_INT
        } catch (e: NumberFormatException) {
            NON_EXISTENT_INT
        }
    }

    override fun toString(): String {
        return "DmeItem(type='$type', vars=$vars, children=$children)"
    }
}